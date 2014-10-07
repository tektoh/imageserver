module.exports = function (grunt) {
  'use strict';

  grunt.initConfig({
    go: {
      options: {
       GOPATH: ["_vendor", "../../../.."]
      },
      imageserver: {
        test_flags: ["-cover"],
      },
      main: {
        run_files: ['main.go']
      }
    },
    watch: {
      files: ['**/*.go'],
      tasks: ['build', 'test']
    }
  });

  grunt.loadNpmTasks('grunt-go');
  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('default', ['watch']);
  grunt.registerTask('build', ['go:build:imageserver']);
  grunt.registerTask('test', ['go:test:imageserver']);
  grunt.registerTask('run', ['build', 'go:run:main']);
};
