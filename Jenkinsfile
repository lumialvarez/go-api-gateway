pipeline {
	agent any
	tools {
        jdk 'JDK'
    }
	environment {
		SSH_MAIN_SERVER = credentials("SSH_MAIN_SERVER")
	
		DATASOURCE_URL_CLEARED = credentials("DATASOURCE_URL_CLEARED")
		DATASOURCE_USERNAME = credentials("DATASOURCE_USERNAME")
		DATASOURCE_PASSWORD = credentials("DATASOURCE_PASSWORD")
	}
	stages {
		stage('Get Version') {
			steps {
				/*script {
					MAVEN_VERSION = sh (
						script: "mvn help:evaluate -Dexpression=project.version -q -DforceStdout",
						returnStdout: true
					).trim()
				}*/
				script {
                    MAVEN_VERSION = "1.0.0"
                }
				script {
					currentBuild.displayName = "#" + currentBuild.number + " - v" + MAVEN_VERSION
				}
			}
		}
		stage('Test') {
			steps {
				//sh 'go test ./...'
				sh 'echo test'
			}
		}
		stage('Deploy') {
			steps {
				script {
					REMOTE_HOME = sh (
						script: "ssh ${SSH_MAIN_SERVER} 'pwd'",
						returnStdout: true
					).trim()
				}
				//script_internal_ip.sh -> ip route | awk '/docker0 /{print $9}'
				script {
					INTERNAL_IP = sh (
						script: "ssh ${SSH_MAIN_SERVER} 'sudo bash script_internal_ip.sh'",
						returnStdout: true
					).trim()
				}


                sh 'java ReplaceSecrets.java DATASOURCE_URL_CLEARED $DATASOURCE_URL_CLEARED'
                sh 'java ReplaceSecrets.java DATASOURCE_USERNAME $DATASOURCE_USERNAME'
                sh 'java ReplaceSecrets.java DATASOURCE_PASSWORD $DATASOURCE_PASSWORD'
                sh 'cat src/cmd/devapi/config/envs/prod.env'


				sh "echo '${BUILD_TAG}' > BUILD_TAG.txt"
				
				sh "ssh ${SSH_MAIN_SERVER} 'sudo rm -rf ${REMOTE_HOME}/tmp_jenkins/${JOB_NAME}'"
				sh "ssh ${SSH_MAIN_SERVER} 'sudo mkdir -p -m 777 ${REMOTE_HOME}/tmp_jenkins/${JOB_NAME}'"
				
				sh "scp -r ${WORKSPACE}/* ${SSH_MAIN_SERVER}:${REMOTE_HOME}/tmp_jenkins/${JOB_NAME}"
				
			
				sh "ssh ${SSH_MAIN_SERVER} 'sudo docker rm -f go-api-gateway &>/dev/null && echo \'Removed old container\''"
				
				sh "ssh ${SSH_MAIN_SERVER} 'cd ${REMOTE_HOME}/tmp_jenkins/${JOB_NAME} ; sudo docker build . -t go-api-gateway'"

				sh "ssh ${SSH_MAIN_SERVER} 'sudo docker run --name go-api-gateway --net=backend-services --add-host=lmalvarez.com:${INTERNAL_IP} -p 9191:9191 -e SCOPE='prod' -d --restart unless-stopped go-api-gateway:latest'"
			}
		}
	}
}