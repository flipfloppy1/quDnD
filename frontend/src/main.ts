import "./style.css";
import "./app.css";

//import logo from "./assets/images/logo-universal.png";
import * as app from "../wailsjs/go/main/App";
import * as cat from "../wailsjs/go/main/Categories";
import { main } from "../wailsjs/go/models";

var screenDefaults = new Map<string, HTMLDivElement>();

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
  console.log(screens);
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

function wipeScreen(screen: string): HTMLDivElement {
  console.log(screen);
  let domScreen = document.getElementById(screen) as HTMLDivElement;
  console.log(domScreen);
  console.log(screenDefaults.get(screen));

  domScreen.remove();
  let content = document.getElementsByClassName("content")[0] as HTMLDivElement;
  content.appendChild(
    screenDefaults.get(screen)?.cloneNode(true) as HTMLDivElement,
  );
  return document.getElementById(screen) as HTMLDivElement;
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
    wipeScreen(category + "Screen");
    let screen = switchScreen(category + "Screen");
    let parent = screen.parentElement;
    if (parent) {
      parent.removeChild(screen);
      parent.appendChild(
        screenDefaults
          .get(category + "Screen")
          ?.cloneNode(true) as HTMLDivElement,
      );
      screen = document.getElementById(category + "Screen") as HTMLDivElement;
    }
    screen.style.display = "block";
    (
      screen.getElementsByClassName("entryTitle")[0] as HTMLHeadingElement
    ).innerText = "Loading...";
    app.GeneratePage(pageid).then((pageInfo: main.PageInfo) => {
      (
        screen.getElementsByClassName("entryTitle")[0] as HTMLHeadingElement
      ).innerText = pageInfo.pageTitle;
      if (pageInfo.description) {
        let desc = document.createElement("p");
        desc.className = "entryDesc";
        desc.innerText = pageInfo.description.trim();
        (
          screen.getElementsByClassName("entryCard")[0] as HTMLDivElement
        ).appendChild(desc);
      }
      if (pageInfo.imgSrc) {
        let entryImg = document.createElement("img");
        entryImg.className = "entryImg";
        entryImg.src = "https://wiki.cavesofqud.com" + pageInfo.imgSrc;
        (
          screen.getElementsByClassName("entryOverview")[0] as HTMLDivElement
        ).appendChild(entryImg);
      }
      if (pageInfo.statblock) {
        // TODO
        let statblockButton = document.createElement("button");
        statblockButton.className = "entryStatblockDropdown";
        let statBtnText = document.createElement("p");
        statBtnText.className = "statblockBtnText";
        statBtnText.innerText = "Minimize statblock";
        statblockButton.appendChild(statBtnText);
        statblockButton.onclick = () => {
          let dropdown = screen.getElementsByClassName(
            "entryStatblockDropdown",
          )[0] as HTMLButtonElement;
          let statblockContent = dropdown.getElementsByClassName(
            "statblockContent",
          )[0] as HTMLDivElement;
          if (statblockContent.style.display == "none") {
            statblockContent.style.display = "flex";
            statBtnText.innerText = "Minimize statblock";
          } else {
            statblockContent.style.display = "none";
            statBtnText.innerText = "Expand statblock";
          }
        };
        screen
          .getElementsByClassName("entryContent")[0]
          .appendChild(statblockButton);

        screen
          .getElementsByClassName("entryStatblockDropdown")[0]
          .insertAdjacentHTML(
            "beforeend",
            `<div class="statblockContent"><div class="statblockAttributes"></div><div class="statblockCoreStats"><div class="statblockAC"><img class="statblockACImg"></img><p class="statblockACText">` +
              pageInfo.statblock.stats[main.Stat.AC] +
              `</p></div><div class="statblockHP"><img class="statblockHPImg"></img><p class="statblockHPText">` +
              pageInfo.statblock.stats[main.Stat.HP] +
              `</p></div></div>`,
          );
      }
      (
        screen.getElementsByClassName("wikiDropdown")[0] as HTMLButtonElement
      ).onclick = () => {
        let elem = screen.getElementsByClassName(
          "wikiDropdown",
        )[0] as HTMLButtonElement;
        console.log("Triggered");
        if (elem.value == "minimized") {
          elem.value = "maximized";
          elem.className = "wikiDropdown wikiActive";
          (elem.lastChild as HTMLIFrameElement).style.display = "block";
          elem.getElementsByTagName("p")[0].innerText = "Minimize wiki page";
          console.log(elem);
        } else {
          elem.value = "minimized";
          elem.className = "wikiDropdown";
          (elem.lastChild as HTMLIFrameElement).style.display = "none";
          elem.getElementsByTagName("p")[0].innerText = "Expand wiki page";
        }
      };
      let iframe = document.createElement("iframe");
      iframe.src =
        "https://wiki.cavesofqud.com/wiki/Special:Redirect/page/" + pageid;
      iframe.style.display = "none";
      iframe.style.width = "100%";
      iframe.style.height = "1000px";
      iframe.className = "wikiFrame";
      (
        screen.getElementsByClassName("wikiDropdown")[0] as HTMLButtonElement
      ).appendChild(iframe);
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

  [].slice
    .call(document.getElementsByClassName("screen"))
    .forEach((elem: HTMLDivElement) => {
      screenDefaults.set(elem.id, elem.cloneNode(true) as HTMLDivElement);
    });

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
