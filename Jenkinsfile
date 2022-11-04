pipeline{
    agent {
        node {
            label "fe"
        }
    }
    tools {
        go 'golang-1.18'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0
    }
    stages {
        stage("Build"){
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'go get ./...'
                sh 'docker build -t registry.gitlab.com/getfit-tech/golang-test .'
            }
        }
    }
}