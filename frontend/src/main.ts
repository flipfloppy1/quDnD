import "./style.css";
import "./app.css";

//import logo from "./assets/images/logo-universal.png";
import * as go from "../wailsjs/go/main/App";
//import * as models from "../wailsjs/go/models";

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

function goToPage(key: string) {
  go.GeneratePage(key);
}

function searchText(query: string) {
  console.log("searchText");
  let searchStatus = document.getElementById("searchStatus");
  if (searchStatus) {
    searchStatus.innerText = "Fetching...";
  }
  go.SearchForPage(query).then((pages: any) => {
    console.log(pages);
    if (searchStatus) {
      if (pages.length) {
        searchStatus.innerText = "";
        let results = document.getElementById(
          "searchResults",
        ) as HTMLDivElement;
        if (results) {
          console.log(pages);
          results.innerHTML = "";
          pages.forEach((page: any) => {
            results.innerHTML = results.innerHTML.concat(
              `<div class="searchResult"><h3 class="searchLink">` +
                page.title +
                `</h3></div>`,
            );
            let result = results.lastChild as HTMLDivElement;
            result.innerHTML = result.innerHTML.concat(page.excerpt);
            let h3 = result.firstChild?.firstChild as HTMLHeadingElement;
            h3.onclick = () => {
              goToPage(page.key);
            };
          });
        }
      } else {
        searchStatus.innerText = "No results found!";
      }
    }
  });
}

function frontendInit() {
  go.LoadPages();
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
