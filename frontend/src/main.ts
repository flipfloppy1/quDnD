import "./style.css";
import "./app.css";

//import logo from "./assets/images/logo-universal.png";
import { SearchForPage } from "../wailsjs/go/main/App";
//import { PageInfo } from "../wailsjs/go/models";

function switchScreen(screen: string) {
  let screens = document.getElementsByClassName("screen");
  for (let i = 0; i < screens.length; i++) {
    let domScreen = screens[i] as HTMLDivElement;
    if (domScreen.id == screen) {
      domScreen.style.display = "block";
    } else {
      domScreen.style.display = "none";
    }
  }
}

function setNavHandlers() {
  let navElems = [].slice.call(
    document.getElementsByClassName("navelement"),
  ) as HTMLButtonElement[];
  navElems.forEach((elem) => {
    elem.onclick = (event) => {
      switchScreen((event.target as HTMLButtonElement).value + "Screen");
      console.log((event.target as HTMLButtonElement).value);
    };
  });
}

function searchText(query: string) {
  console.log("searchText");
  let searchStatus = document.getElementById("searchStatus");
  if (searchStatus) {
    searchStatus.innerText = "Fetching...";
  }
  SearchForPage(query).then((page: string) => {
    console.log(page);
    if (searchStatus) {
      if (page) {
        searchStatus.innerText = "";
        let results = document.getElementById("searchResult");
        if (results) {
          results.innerHTML = page;
        }
      } else {
        searchStatus.innerText = "No results found!";
      }
    }
  });
}

function frontendInit() {
  switchScreen("searchScreen");
  setNavHandlers();

  let searchBox = document.getElementById("searchBox");
  if (searchBox) {
    searchBox.addEventListener("keyup", (event) => {
      if (event.key == "Enter") {
        searchText((event.target as HTMLInputElement).value);
      }
    });
  }
}

frontendInit();
