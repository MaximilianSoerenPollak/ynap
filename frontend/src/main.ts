import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
// import {Greet} from '../wailsjs/go/main/App';
import {MakeAccount} from '../wailsjs/go/main/App';
import {GetAllAccounts} from '../wailsjs/go/main/App';

// Setup the greet function
// window.greet = function () {
//     // Get name
//     let name = nameElement!.value;
//
//     // Check if the input is empty
//     if (name === "") return;
//
//     // Call App.Greet(name)
//     try {
//         Greet(name)
//             .then((result) => {
//                 // Update result with data back from App.Greet()
//                 resultElement!.innerText = result;
//             })
//             .catch((err) => {
//                 console.error(err);
//             });
//     } catch (err) {
//         console.error(err);
//     }
// };
window.makeAccount = function() {
    // Get firstname 
    let firstName = firstNameElement!.value;
    let lastName = lastNameElement!.value;
    if (firstName == "")
        return "please enter a first name"
    if (lastName == "")
        return "please enter a last name"
    try {
        MakeAccount(firstName, lastName)
            .then((result) => {
                response_accountcreation!.innerText = result; 
            })
            .catch((err) => {
                console.error(err);
            });
    }catch (err) {
        console.error(err);
    }
    return
}

window.getAllAccounts = function() {
    try {
        GetAllAccounts()
        .then((result) => {
            responseTable!.textContent = JSON.stringify(result, null , 4)

        })
        .catch((err: any) => {
            console.error(err);
        });
    }catch (err) {
        console.log("Inside the error catch")
        console.error(err);
    }
    console.log("Ran through the function 'Get all accounts")
    return
}
function insertAccountsOnSelection() {
    GetAllAccounts().then((result) => {
        let accSelector = document.getElementById("account-selector") as HTMLSelectElement;
        result.forEach( (acc) => {
            let option = document.createElement("option");
            option.text = acc.first_name!
            option.value = acc.first_name!
            accSelector.add(option)
        })})
        .catch((err: any) => {
            console.error(err);
        }); 
}

// <li id="accountname_response"></li>
// <li id="firstname_response"></li>
// <li id="lastname_response"></li>
// <li id="balance_response"></li>
// <li id="portfolios_response"></li>
// <li id="updatedat_response"></li>
// <li id="createdat_response"></li>
//
function displayCurrentAccount() {
    let currAcc =  currentSelectedAccount!.value
    let currAccDisplay = document.getElementById("selected-account")
    currAccDisplay!.innerText = currAcc
}

document.querySelector('#app')!.innerHTML = `
        <img id="logo" class="logo">
        <pre id="responsetablepre"></pre>
        <button class="btn" onclick="insertAccountsOnSelection()">Select accounts</button>
        <div id="test"></div>
        <div id="account-selector-div">
            <label for 'account-selector'>"Select the active account here"</label>
            <h3 id="selected-acccount"></h3>
            <select name="account-selector" id="account-selector" onchange="displayCurrentAccount()"></select>
        </div>
        <div class="input-box", "accountcreation" id="input">
            <h3 id="response-accountcreation">"CinnerTextacccount here"</h3>
            <input class="input" id="firstName" type="text" autocomplete="off" placeholder="Enter your first name here"/>
            <input class="input" id="lastName" type="text" autocomplete="off" placeholder="Enter your last name here" />
        </div>
        <div>
            <button class="btn" onclick="makeAccount()" onclick="insertAccountsOnSelection()" >Create Account</button>
        </div>
        <div class="input-box", "portfoliocreation" id="input">
            <h3 id="response-accountcreation">"Create a portfoolio here"</h3>
            <input class="input" id="accountid" type="text" autocomplete="off" placeholder="Enter your first name here"/>
            <input class="input" id="portfolio" type="text" autocomplete="off" placeholder="Enter your last name here" />
        </div>
        <div>
            <button class="btn" onclick="getAllAccounts()">Get all accounts"</button>
        </div>
`;

(document.getElementById('logo') as HTMLImageElement).src = logo;

// let nameElement = (document.getElementById("name") as HTMLInputElement);
// nameElement.focus();
let firstNameElement = (document.getElementById("firstName") as HTMLInputElement);
firstNameElement.focus();
let lastNameElement = (document.getElementById("lastName") as HTMLInputElement);
lastNameElement.focus();
let response_accountcreation = document.getElementById("response-accountcreation");
let responseTable = document.getElementById("responsetablepre");
let currentSelectedAccount = document.getElementById("account-selector") as HTMLSelectElement;

declare global {
    interface Window {
        greet: () => void;
        makeAccount: () => void;
        getAllAccounts: () => void;
    }
}
