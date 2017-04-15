pipeline {
  agent any
  stages {
    stage('build something') {
      steps {
        input(message: 'Do you want to build?', id: 'foo', ok: 'bar')
      }
    }
    stage('moo') {
      steps {
        milestone(ordinal: 1, label: '11')
      }
    }
  }
}