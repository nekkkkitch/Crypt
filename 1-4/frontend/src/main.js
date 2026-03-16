import './style.css';
import './app.css';

import {ChangeCypher} from '../wailsjs/go/app/App';
import {ChangeParams} from '../wailsjs/go/app/App';
import {Cypher} from '../wailsjs/go/app/App';
import {Decypher} from '../wailsjs/go/app/App';

var cypherField = document.getElementById("cypher-input");
var decypherField = document.getElementById("decypher-input");
var resultElement = document.getElementById("result");
var paramsField = document.getElementById("param1");
// Setup the greet function
window.cypher = function () {
    let input = cypherField.value;

    // Call App.Greet(name)
    try {
        Cypher(input)
            .then((result) => {
                // Update result with data back from App.Greet()
                resultElement.value = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

window.decypher = function () {
    let input = decypherField.value;

    // Call App.Greet(name)
    try {
        Decypher(input)
            .then((result) => {
                // Update result with data back from App.Greet()
                resultElement.value = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

window.changeCypher = function () {
    let select = document.getElementById("cypher-select").value;

    try {
        ChangeCypher(select);
        switch (select) {
            case "atbash":
                document.getElementById("param-label").innerHTML = "no param :(";
                paramsField.style.display = "none";
                break;
            case "scytale":
                document.getElementById("param-label").innerHTML = "Height";
                paramsField.style.display = "inline ";
                break;
            case "polybius":
                document.getElementById("param-label").innerHTML = "Language (0 - English, 1 - Russian)";
                paramsField.style.display = "inline";
                break;
            case "caesar":
                document.getElementById("param-label").innerHTML = "Step";
                paramsField.style.display = "inline";
                break;
        }

    } catch (err) {
        console.error(err);
    }
}

window.changeParams = function (value) {
    let param1 = Number(value);
    
    try { 
        ChangeParams(Number(value));
    } catch (err) {
        console.error(err);
    }
}