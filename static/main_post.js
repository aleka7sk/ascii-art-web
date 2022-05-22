result = document.getElementById("result_api")

fetch('http://localhost:8080/api')
    .then(result => result.json())
    .then((output) => {
       result.textContent = output.Text;
        
}).catch(err => console.error(err));

function clearOutput() {
    output = document.getElementsByClassName("output")
    for(let i = 0; i < output.length; i++) {
        output[i].textContent = ""
    }
}

