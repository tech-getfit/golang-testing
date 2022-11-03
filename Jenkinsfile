pipeline{
    agent {
        node {
            label "fe"
        }
    }
    stages {
        stage("Hello"){
            steps{
                sh 'go version'
            }
        }
    }
}