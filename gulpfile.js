var gulp = require('gulp');
var $    = require('gulp-load-plugins')();
var sourcemaps = require('gulp-sourcemaps');

var sassPaths = [
  'node_modules/bootstrap/scss'
];

gulp.task('sass', function() {
  return gulp.src('wwwroot/scss/app.scss')
    .pipe($.sourcemaps.init())
    .pipe($.sass({
      includePaths: sassPaths,
      outputStyle: 'compressed' // if css compressed **file size**
    })
      .on('error', $.sass.logError))
    .pipe($.autoprefixer({
      browsers: ['last 2 versions', 'ie >= 9']
    }))
    .pipe(sourcemaps.write())
    .pipe(gulp.dest('wwwroot/public/css'));
});

gulp.task('default', ['sass'], function() {
  gulp.watch(['wwwroot/scss/**/*.scss'], ['sass']);
});
