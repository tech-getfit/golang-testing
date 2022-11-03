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
    }
    stages {
        stage("Hello"){
            steps{
                sh 'go version'
            }
        }
    }
}