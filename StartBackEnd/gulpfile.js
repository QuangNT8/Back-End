"use strict";

var gulp = require('gulp');

var path = require('path');

var bower = require('gulp-bower');
var exit = require('gulp-exit');
var gprocess = require('gulp-process');
var shell = require('gulp-shell');
var jasmineBrowser = require('gulp-jasmine-browser');
var webpack = require('webpack-stream');

gulp.task('bower', function () {
  return bower();
});

// Build gRPC-Server
gulp.task('server', shell.task([
  'go build -o bin/grpc-server back-end/cmd/cmd-grpc-server',
]));

// Build Gateway-Server
gulp.task('gateway', shell.task([
  'go build -o bin/gateway-server back-end/cmd/cmd-gateway-server',
]));

gulp.task('grpc-server', ['server'], function () {
  gprocess.start('server-server', 'bin/grpc-server', [
    '--logtostderr',
  ]);
  gulp.watch('bin/grpc-server', ['grpc-server']);
});

gulp.task('serve-gateway', ['gateway', 'grpc-server'], function () {
  gprocess.start('gateway-server', 'bin/gateway-server', [
    '--logtostderr', '--openapi_dir', path.join(__dirname, "../proto"),
  ]);
  gulp.watch('bin/gateway-server', ['serve-gateway']);
});

gulp.task('backends', ['serve-gateway', 'grpc-server']);

var specFiles = ['*.spec.js'];
gulp.task('test', ['backends'], function (done) {
  return gulp.src(specFiles)
    .pipe(webpack({ output: { filename: 'spec.js' } }))
    .pipe(jasmineBrowser.specRunner({
      console: true,
      sourceMappedStacktrace: true,
    }))
    .pipe(jasmineBrowser.headless({
      findOpenPort: true,
      catch: true,
      throwFailures: true,
    }))
    .on('error', function (err) {
      done(err);
      process.exit(1);
    })
    .pipe(exit());
});

gulp.task('serve', ['backends'], function (done) {
  var JasminePlugin = require('gulp-jasmine-browser/webpack/jasmine-plugin');
  var plugin = new JasminePlugin();

  return gulp.src(specFiles)
    .pipe(webpack({
      output: { filename: 'spec.js' },
      watch: true,
      plugins: [plugin],
    }))
    .pipe(jasmineBrowser.specRunner({
      sourceMappedStacktrace: true,
    }))
    .pipe(jasmineBrowser.server({
      port: 8000,
      whenReady: plugin.whenReady,
    }));
});

