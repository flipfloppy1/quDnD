import "./style.css";
import "./app.css";

//import logo from "./assets/images/logo-universal.png";
import * as app from "../wailsjs/go/main/App";
import * as cat from "../wailsjs/go/main/Categories";
import { main } from "../wailsjs/go/models";

function switchScreen(screen: string): HTMLDivElement {
  let nav = document.getElementById(screen.replace("Screen", "Nav"));
  if (nav) {
    [].slice
      .call(nav.parentElement?.children)
      .forEach((navElem: HTMLButtonElement) => {
        navElem.className = "navelement";
      });
    nav.className = "navelement navactive";
  }
  let screens = [].slice.call(document.getElementsByClassName("screen"));
  for (let i = 0; i < screens.length; i++) {
    let domScreen = screens[i] as HTMLDivElement;
    if (domScreen.id == screen) {
      domScreen.style.display = "block";
    } else {
      domScreen.style.display = "none";
    }
  }

  return screens.filter((val: HTMLDivElement) => {
    return val.id === screen;
  })[0];
}

function setNavHandlers() {
  let navElems = [].slice.call(
    document.getElementsByClassName("navelement"),
  ) as HTMLButtonElement[];
  navElems.forEach((elem) => {
    elem.onclick = (event) => {
      let elem = event.target as HTMLButtonElement;
      switchScreen(elem.value + "Screen");
    };
  });
}

function goToPage(pageid: number) {
  cat.GetScreen(pageid).then((category: main.Screen) => {
    let screen = switchScreen(category + "Screen");
    (
      screen.getElementsByClassName("entryTitle")[0] as HTMLHeadingElement
    ).innerText = "Loading...";
    app.GeneratePage(pageid).then((pageInfo: main.PageInfo) => {
      (
        screen.getElementsByClassName("entryTitle")[0] as HTMLHeadingElement
      ).innerText = pageInfo.pageTitle;
    });
  });
}

function searchText(query: string) {
  console.log("searchText");
  let searchStatus = document.getElementById("searchStatus");
  if (searchStatus) {
    searchStatus.innerText = "Fetching...";
  }
  app.SearchForPage(query).then((pages: main.RestPageSearch) => {
    console.log(pages);
    let pageResults = pages.pages;
    if (searchStatus) {
      if (pageResults.length) {
        searchStatus.innerText = "";
        let results = document.getElementById(
          "searchResults",
        ) as HTMLDivElement;
        if (results) {
          console.log(pages);
          results.innerHTML = "";
          pageResults.forEach((page: any) => {
            results.innerHTML = results.innerHTML.concat(
              `<div class="searchResult"><h3 class="searchLink">` +
                page.title +
                `</h3></div>`,
            );
            let result = results.lastChild as HTMLDivElement;
            result.innerHTML = result.innerHTML.concat(page.excerpt);
          });
          setTimeout(() => {
            [].slice
              .call(results.children)
              .forEach((child: HTMLElement, i: number) => {
                let h3 = child.firstChild as HTMLElement;
                console.log(h3);
                h3.onclick = () => {
                  console.log(pageResults[i].id);
                  goToPage(pageResults[i].id);
                };
              });
          }, 200);
        }
      } else {
        searchStatus.innerText = "No results found!";
      }
    }
  });
}

function frontendInit() {
  cat.LoadCategories();
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
