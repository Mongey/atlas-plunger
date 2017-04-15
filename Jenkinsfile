pipeline {
  agent any
  stages {
    stage('build something') {
      steps {
        input(message: 'Do you want to build?', id: 'foo', ok: 'Deploy to UAT', submitter: 'what\'s this', submitterParameter: 'smsms')
      }
    }
    stage('moo') {
      steps {
        milestone(ordinal: 1, label: '11')
      }
    }
  }
}