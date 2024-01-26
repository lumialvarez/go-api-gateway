def APP_VERSION
pipeline {
   agent any
   stages {
      stage('Get Version') {
         steps {
            script {
               APP_VERSION = sh (
                  script: "grep -m 1 -Po '[0-9]+[.][0-9]+[.][0-9]+' CHANGELOG.md",
                  returnStdout: true
               ).trim()
            }
            script {
               currentBuild.displayName = "#" + currentBuild.number + " - v" + APP_VERSION
            }
            script{
                if(currentBuild.previousSuccessfulBuild) {
                    lastBuild = currentBuild.previousSuccessfulBuild.displayName.replaceFirst(/^#[0-9]+ - v/, "")
                    echo "Last success version: ${lastBuild} \nNew version to deploy: ${APP_VERSION}"
                    if(lastBuild == APP_VERSION)  {
                         currentBuild.result = 'ABORTED'
                         error("Aborted: A version that already exists cannot be deployed a second time")
                    }
                }
            }
         }
      }
   }
}