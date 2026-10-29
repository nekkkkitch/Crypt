import './style.css';
import './app.css';

import { Analyse } from '../wailsjs/go/app/App';
import { SetText } from '../wailsjs/go/app/App';
import { SetFile } from '../wailsjs/go/app/App';

import Highcharts from 'highcharts';
import 'highcharts/modules/exporting';
import 'highcharts/modules/accessibility';


window.setFile = function () {
    try{
        let bool = document.querySelector('input[name="language"]:checked').value;
        SetFile(bool)
    } catch(err){
        console.error(err)
    }
};

window.setText = function () {
    try{
        let text = document.getElementById("input-text").value
        let bool = document.querySelector('input[name="language"]:checked').value;
        SetText(text, bool)
    }catch(err){
        console.error(err)
    }
}

window.analyse = function () {
    try{
        Analyse().then((result)=>{
            console.log(result)
            let fr = result["freq_data"]["frequencies"]
            let k = result["k_data"]
            let ac = result["ac_data"]
            let ind = result["index_data"]
            console.log(fr)

            let symbols = []
            let data = []
            for (let key of Object.keys(fr)) {
                var value = fr[key];
                symbols.push(key)
                data.push(value["part"])
            }

            Highcharts.chart('container_fr', {
                title:{
                    text: 'Frequency Analysis' 
                },
                xAxis:{
                    categories: symbols,
                    crosshair: true,
                    accessibility: {
                        description: 'Symbols'
                    }   
                },
                yAxis:[
                    {
                    title: { text: 'Frequency' },
                    }
                ],
                series: [{
                    name: 'Histogram',
                    type: 'column',
                    data,
                    zIndex: 1,
                    opacity: 0.9
                }],
                plotOptions: {
                    column: {
                        pointPadding: 0,
                        borderWidth: 0,
                        groupPadding: 0,
                        shadow: false
                    }
                }
            });

            let coin = document.getElementById("coincidence")
            coin.innerHTML = "Coincidence probability = " + (ind*100).toFixed(2) + "%"

            console.log(Object.keys(k["possible_dividers"]))
            let key = document.getElementById("key-lengths")
            key.innerHTML = "Possible key lengths: " + Object.keys(k["possible_dividers"])

            let t = []
            let amounts = []
            for (var i = 0; i < ac["correlations"].length; i++) {
                var value = ac["correlations"][i];
                t.push((i+1).toString())
                amounts.push(parseInt(value["amount"]))
            }

            console.log("t", t)
            console.log("amounts", amounts)

            Highcharts.chart('container_ac', {
                title:{
                    text: 'AutoCorrelation Analysis' 
                },
                xAxis:{
                    categories: t,
                    crosshair: true,
                    accessibility: {
                        description: 'Move'
                    }   
                },
                yAxis:{
                    title: { text: 'Correlations' },
                },
                series: [{
                    name: 'Correlation',
                    type: 'column',
                    data: amounts,
                    zIndex: 1,
                    opacity: 0.9
                }],
                plotOptions: {
                    column: {
                        pointPadding: 0,
                        borderWidth: 0,
                        groupPadding: 0,
                        shadow: false
                    }
                }
            });            
        })
    }catch(err){
        console.error(err)
    }
}
