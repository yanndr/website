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

gulp.task('vendors',function(done){
  console.log('Copying jquery');
    gulp.src([
      './node_modules/jquery/dist/*.js'
    ]).pipe(gulp.dest('./wwwroot/public/js/vendors/jquery/'));

    console.log('Copying bootstrap');
    gulp.src([
      './node_modules/bootstrap/dist/css/*.*'
    ]).pipe(gulp.dest('./wwwroot/public/css/vendors/bootstrap/'));

    gulp.src([
      './node_modules/bootstrap/dist/js/*.*'
    ]).pipe(gulp.dest('./wwwroot/public/js/vendors/bootstrap/'));

    console.log('Copying popper.js');
    gulp.src([
      './node_modules/popper.js/dist/*.*'
    ]).pipe(gulp.dest('./wwwroot/public/js/vendors/popper.js/'));
    
    console.log('Copying scrollme.js');
    gulp.src([
      './wwwroot/vendors/*.js'
    ]).pipe(gulp.dest('./wwwroot/public/js/vendors/'));

})
