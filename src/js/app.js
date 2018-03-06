
var $animation_elements = $('.animate');
var $changing_elements = $('.change');
var $window = $(window);
var sdegree = 0;

function check_if_in_view() {
  var window_height = $window.height();
  var window_top_position = $window.scrollTop();
  var window_bottom_position = (window_top_position + window_height);
  
  if(window_top_position > 600){
    $changing_elements.addClass("changed")
  }else{
    $changing_elements.removeClass("changed")
  }
  $.each($animation_elements, function() {
    var $element = $(this);
    var element_height = $element.outerHeight();
    var element_top_position = $element.offset().top;
    var element_bottom_position = (element_top_position + element_height);
    // console.log(window_bottom_position+"-"+ window_top_position)
    //check to see if this current container is within viewport
    if ((element_bottom_position >= window_top_position) &&
      (element_top_position <= window_bottom_position)) {
      $element.addClass('in-view');
      $element.addClass('viewed');
    } else {
      $element.removeClass('in-view');
    }
    if (element_top_position <  window_top_position+100){
      $element.addClass('on-top');
    }
    else{
      $element.removeClass('on-top');
    }
    if (element_top_position >  window_top_position-100){
      $element.addClass('on-bottom');
    }
    else{
      $element.removeClass('on-bottom');
    }
    if ((element_top_position <  window_bottom_position-100) && (element_top_position >  window_top_position+100)){
       
      $element.addClass('on-middle');
    }
    else{
      $element.removeClass('on-middle');
    }
  });
}


$(window).on('scroll resize', check_if_in_view);
$(window).scroll();
