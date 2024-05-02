
function processing () {

    var url = document.getElementById('urlInput').value
    var searchFormat = document.getElementById('searchFormat').value
    // var filter = document.getElementById('filterInput').value

    var limit = document.getElementById('limitNumber').value
    
    if (limit != "") {
        limit = Number(limit)
    } else {
        limit = 0
    }
    
    fetch('http://localhost:3000/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url, searchFormat, limit })
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


document.getElementById('limit').addEventListener('change', (e) => {
    var limitInput = document.getElementById("limitInput")
    var selectedOption = e.target.value
    if (selectedOption == "yes") {
        limitInput.style.display = "block";
    } else {
        limitInput.style.display = "none";
    }
})

document.getElementById('myForm').addEventListener('submit', (e) => {
    e.preventDefault(); 
    processing();
})


