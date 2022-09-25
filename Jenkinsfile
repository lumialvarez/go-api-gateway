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

		DOCKERHUB_CREDENTIALS=credentials('dockerhub-lmalvarez')
	}
	stages {
		stage('Get Version') {
			steps {
				script {
					APP_VERSION = sh (
						script: '''grep -m 1 -Po '[0-9]+[.][0-9]+[.][0-9]+' CHANGELOG.md ''',
						returnStdout: true
					).trim()
				}
				script {
					currentBuild.displayName = "#" + currentBuild.number + " - v" + APP_VERSION
				}
			}
		}
		stage('Test') {
			steps {
				//sh 'go test ./...'
				sh 'echo test'
			}
		}
		stage('Build') {
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

                sh '''echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin '''

                sh '''ssh ${SSH_MAIN_SERVER} 'sudo rm -rf ${REMOTE_HOME}/tmp_jenkins/${JOB_NAME}' '''
                sh '''ssh ${SSH_MAIN_SERVER} 'sudo mkdir -p -m 777 ${REMOTE_HOME}/tmp_jenkins/${JOB_NAME}' '''

                sh '''scp -r ${WORKSPACE}/* ${SSH_MAIN_SERVER}:${REMOTE_HOME}/tmp_jenkins/${JOB_NAME} '''

                sh '''ssh ${SSH_MAIN_SERVER} 'sudo docker rm -f lmalvarez/go-api-gateway:${APP_VERSION} &>/dev/null && echo \'Removed old container\'' '''

                sh '''ssh ${SSH_MAIN_SERVER} 'cd ${REMOTE_HOME}/tmp_jenkins/${JOB_NAME} ; sudo docker build . -t lmalvarez/go-api-gateway:${APP_VERSION}' '''

                sh '''ssh sudo docker push lmalvarez/go-api-gateway:${APP_VERSION}' '''
            }
        }
		stage('Deploy') {
			steps {
				sh '''ssh ${SSH_MAIN_SERVER} 'sudo docker run --name go-api-gateway --net=backend-services --add-host=lmalvarez.com:${INTERNAL_IP} -p 9191:9191 -e SCOPE='prod' -d --restart unless-stopped lmalvarez/go-api-gateway:${APP_VERSION}' '''
			}
		}
	}
}