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
                resultElement!.innerText = result; 
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
            console.log(result)
            resultElement!.innerText = result;
        })
        .catch((err: any) => {
            console.error(err);
        });
    }catch (err) {
        console.error(err);
    }
    return
}

document.querySelector('#app')!.innerHTML = `
    <img id="logo" class="logo">
      <div class="result" id="result">Please enter your first Name</div>
      <div class="input-box" id="input">
        <input class="input" id="firstName" type="text" autocomplete="off" />
      </div>
      <div class="result" id="result">Please enter your last Name</div>
      <div class="input-box" id="input">
        <input class="input" id="lastName" type="text" autocomplete="off" />
        </div>
        <button class="btn" onclick="makeAccount()">Create Account</button>
        </div>
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
let resultElement = document.getElementById("result");

declare global {
    interface Window {
        greet: () => void;
        makeAccount: () => void;
        getAllAccounts: () => void;
    }
}
