
function processing () {

    var url = document.getElementById('urlInput').value
    let searchFormat = document.getElementById('searchFormatInput').value
    let filter = document.getElementById('filterInput').value


    console.log("URL: ", url);
    console.log("searchFormat: ", searchFormat);
    console.log("filter: ", filter);

    fetch('http://localhost:3000/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url, filter, searchFormat })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.text();
    })
    .then(data => {
        console.log('Response from server: ', data);
    })
    .catch(error => {
        console.log('error: ', error);
    })


}


document.getElementById('myForm').addEventListener('submit', (e) => {
    e.preventDefault(); 
    processing();
})


