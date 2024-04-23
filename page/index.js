
function processing () {

    const url = document.getElementById('urlInput').value
    var type = document.getElementById('typeInput').value
    var filter = document.getElementById('filterInput').value


    console.log("URL: ", url);
    console.log("type: ", type);
    console.log("filter: ", filter);


}


document.getElementById('myForm').addEventListener('submit', (e) => {
    e.preventDefault(); 
    processing();
})