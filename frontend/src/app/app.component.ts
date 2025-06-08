import { Component } from "@angular/core";
import { DomSanitizer, SafeResourceUrl } from "@angular/platform-browser";
import { MatSidenavModule } from "@angular/material/sidenav";
import { SidenavComponent } from "./sidenav/sidenav.component";
import { SearchPageComponent } from "./search-page/search-page.component";
import { StatblockComponent } from "./statblock/statblock.component";
import { CommonModule } from "@angular/common";
import { MatButtonModule } from "@angular/material/button";
import { MatInputModule } from "@angular/material/input";
import { MatIconModule } from "@angular/material/icon";
import { MatTooltipModule } from "@angular/material/tooltip";
import { MatProgressSpinnerModule } from "@angular/material/progress-spinner";
import { MatExpansionModule } from "@angular/material/expansion";
import { main } from "../../wailsjs/go/models";
import * as cat from "../../wailsjs/go/main/Categories";
import * as app from "../../wailsjs/go/main/App";

interface SearchPage {
  query: string;
}

@Component({
  selector: "app-root",
  imports: [
    MatSidenavModule,
    SidenavComponent,
    SearchPageComponent,
    CommonModule,
    MatButtonModule,
    MatIconModule,
    MatInputModule,
    MatTooltipModule,
    MatProgressSpinnerModule,
    MatExpansionModule,
    StatblockComponent,
  ],
  templateUrl: "./app.component.html",
  styleUrl: "./app.component.css",
})
export class AppComponent {
  constructor(sanitizer: DomSanitizer) {
    this.sanitizer = sanitizer;
  }
  sanitizer: DomSanitizer;
  title: string = "quDnD";
  name: string = "";
  iframeUrl: string = "";
  category: main.Screen = main.Screen.SEARCH;
  navOpened: boolean = true;
  openedPages: Map<main.Screen, main.PageInfo> = new Map();
  currPage: main.PageInfo | SearchPage | main.Screen = { query: "" };
  loadingPage: boolean = false;

  ngOnInit() {
    document.addEventListener("keydown", (event) => {
      if (event.key === "Escape") {
        this.toggleSidenav();
      }
    });
    this.frontendInit();
  }

  frontendInit() {
    cat.LoadCategories();
  }

  isSearchPage(): boolean {
    return this.category === "search";
  }

  getScreenName(screen: main.Screen): string {
    return screen[0].toUpperCase() + screen.slice(1);
  }

  isEmptyPage(): boolean {
    return typeof this.currPage !== "object";
  }

  goToPage(pageid: number) {
    cat.GetScreen(pageid).then((screen) => {
      this.category = screen;
      this.loadingPage = true;
      this.iframeUrl =
        "https://wiki.cavesofqud.com/Special:Redirect/page/" + pageid;
      app.GeneratePage(pageid).then((page) => {
        this.openedPages.set(screen, page);
        this.currPage = page;
        console.log(page.statblock);
        this.name = page.pageTitle;
        this.loadingPage = false;
        setTimeout(() => {
          let iframe = document.getElementsByClassName(
            "referencePage",
          )[0] as HTMLIFrameElement;
          iframe.src = this.iframeUrl;
        }, 100);
      });
    });
  }

  articleHasDesc(): boolean {
    if (typeof this.currPage === "object")
      return Boolean((this.currPage as main.PageInfo).description);
    else return false;
  }

  articleHasImg(): boolean {
    if (typeof this.currPage === "object")
      return Boolean((this.currPage as main.PageInfo).imgSrc);
    else return false;
  }

  getDescription(): string {
    if (typeof this.currPage === "object") {
      return (this.currPage as main.PageInfo).description ?? "";
    }
    return "";
  }

  getPageInfo(): main.PageInfo | undefined {
    if (typeof this.currPage === "object") {
      return this.currPage as main.PageInfo;
    }

    return undefined;
  }

  getStatblock(): main.Statblock {
    if (typeof this.currPage === "object") {
      let currPage = this.currPage as main.PageInfo;
      if (currPage.statblock) {
        return currPage.statblock;
      }
    }

    return new main.Statblock();
  }

  getIframeUrl(): SafeResourceUrl {
    return this.sanitizer.bypassSecurityTrustResourceUrl(this.iframeUrl);
  }

  getImgSrc(): string {
    if (typeof this.currPage === "object") {
      return (this.currPage as main.PageInfo).imgSrc ?? "";
    }
    return "";
  }

  toggleSidenav() {
    this.navOpened = !this.navOpened;
  }

  screenChange() {
    let catPage = this.openedPages.get(this.category);
    if (catPage) {
      this.currPage = catPage;
      this.name = catPage.pageTitle;
    } else {
      this.name = "";
      this.currPage = this.category;
    }
    this.navOpened = false;
  }

  createCustom() {}
}
