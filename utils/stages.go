package utils

import "fmt"

var SetupStage = CommandsStage{
	SpinnerSuccessMessage: "VPS updated and setup successfully",
	SpinnerFailMessage:    "Error happened running basic setup commands",
	Commands: []string{
		"sudo apt-get update -y",
		"sudo apt-get upgrade -y",
		"sudo apt-get install ca-certificates curl vim",
	},
}

var DockerStage = CommandsStage{
	SpinnerSuccessMessage: "Docker setup successfully",
	SpinnerFailMessage:    "Error happened during setting up docker",
	Commands: []string{
		"sudo apt-get update -y",
		"sudo apt-get install curl -y",
		"sudo apt-get install ca-certificates -y",
		"sudo install -m 0755 -d /etc/apt/keyrings",
		"sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc",
		"sudo chmod a+r /etc/apt/keyrings/docker.asc",
		`echo \
		"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
		$(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
		sudo tee /etc/apt/sources.list.d/docker.list > /dev/null`,
		"sudo apt-get update -y",
		"sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y",
	},
}

func GetTraefikStage(server string) CommandsStage {
	return CommandsStage{
		SpinnerSuccessMessage: "Successfully setup Traefik",
		SpinnerFailMessage:    "Something went wrong setting up Traefik on your VPS",
		Commands: []string{
			"sudo apt-get install git -y",
			"git clone https://github.com/ms-mousa/sidekick-traefik.git",
			fmt.Sprintf(`cd sidekick-traefik && sed -i.bak "s/\$HOST/%s/g; s/\$PORT/%s/g" docker-compose.traefik.yml && rm docker-compose.traefik.yml.bak`, server, "8000"),
			"docker network create sidekick",
			"cd sidekick-traefik && docker compose -p sidekick -f docker-compose.traefik.yml up -d",
		},
	}
}