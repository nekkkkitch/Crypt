import './style.css';
import './app.css';

import { ChangeCypher } from '../wailsjs/go/app/App';
import { ChangeParams } from '../wailsjs/go/app/App';
import { Cypher } from '../wailsjs/go/app/App';
import { Decypher } from '../wailsjs/go/app/App';

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
    console.log("current selection", select)
    try{
        ChangeCypher(select);
        console.log("cypher changed")
    } catch (err){
        console.error(err);
    }
    setTimeout(function(){
        try {
        console.log("sending params")
        switch (select) {
            case "atbash":
                document.getElementById("param-label").innerHTML = "no param :(";
                paramsField.style.display = "none";
                break;
            case "scytale":
                document.getElementById("param-label").innerHTML = "Height";
                paramsField.style.display = "inline";
                paramsField.value = 4;
                changeParams(4);
                break;
            case "polybius":
                document.getElementById("param-label").innerHTML = "Language (0 - English, 1 - Russian)";
                paramsField.style.display = "inline";
                paramsField.value = 0;
                changeParams(0);
                break;
            case "caesar":
                document.getElementById("param-label").innerHTML = "Step";
                paramsField.style.display = "inline";
                paramsField.value = 3;
                changeParams(3);
                break;
            case "gronsfeld":
                document.getElementById("param-label").innerHTML = "Key (digits)";
                paramsField.style.display = "inline";
                paramsField.value = "123";
                changeParams("123");
                break;
            case "vigener":
                document.getElementById("param-label").innerHTML = "Key (letters)";
                paramsField.style.display = "inline";
                paramsField.value = "KEY";
                changeParams("KEY");
                break;
        }
    } catch (err) {
        console.error(err);
    }
    }, 100);
}

window.changeParams = function (value) {
    let param1 = [value];

    try {
        ChangeParams(param1);
    } catch (err) {
        console.error(err);
    }
}