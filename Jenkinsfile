pipeline{
    agent {
        node {
            label "stg"
        }
    }
    tools {
        go 'golang-1.18'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0
        UID_REPO = credentials("uid-getfit-gitlab")
    }
    stages {
        stage("Build"){
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'go get ./...'
                echo("username: ${UID_REPO_USR}")
                echo("password: ${UID_REPO_PSW}")
                // sh 'docker build -t registry.gitlab.com/getfit-tech/golang-test .'
            }
        }
    }
}