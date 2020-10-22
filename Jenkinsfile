pipeline {
    agent { label 'slave' }

    options {
      disableConcurrentBuilds()
      buildDiscarder(logRotator(numToKeepStr: '10'))
    }

    stages {
        stage('Test') {
            agent {
                docker {
                    image 'golang:1.15'
                    reuseNode true
                }
            }

            steps {
                script {
                    sh '''
                        export GOCACHE=/tmp/.cache
                        go test -race -coverprofile=coverage.out -json ./... > test-report.json
                    '''
                }
            }
        }

        stage('Code Scan') {
            steps{
                script{
                    scannerHome = tool 'SonarQubeScanner'
                    withSonarQubeEnv('sonar-config') {
                        sh "${scannerHome}/bin/sonar-scanner"
                    }

                    timeout(time: 2, unit: 'MINUTES') {
                        def qg = waitForQualityGate()
                        if (qg.status != 'OK') {
                            echo 'Unquality Coding'
                        }
                    }
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.withRegistry("https://asia.gcr.io", "gcr:innovationlab-devops") {
                        def image = docker.build("innovationlab-devops/demo-go-pipeline:latest")
                        image.push()
                    }
                }
            }
        }

        stage('K8S Deploy') {
            steps {
                script {
                    withCredentials([file(credentialsId: 'helm-deployer', variable: 'serviceaccountkey')]) {
                        sh "gcloud auth activate-service-account gke-sa@innovationlab-devops.iam.gserviceaccount.com --key-file=${serviceaccountkey}"
                        sh "gcloud container clusters get-credentials sandbox-boostcamp --region asia-southeast1 --project innovationlab-devops"
                        sh 'kubectl apply -f k8s/'
                    }
                }
            }
        }
    }

    post {
        always {
            deleteDir()
        }

        success {
            echo 'Pipeline success'
        }

        failure {
          echo "Pipeline failure"
        }

        aborted {
            echo "Pipeline aborted"
        }
    }
}
