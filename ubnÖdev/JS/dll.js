document.addEventListener('DOMContentLoaded', function() {
    var link = document.getElementById('link');
    // onClick's logic below:
    link.addEventListener('click', function() {

        function searchGoogle() {
        var one = document.getElementById("input1").value;
        var two = 'https://www.google.com/search?q=' + one;
        window.location = two;
    }


    });
});