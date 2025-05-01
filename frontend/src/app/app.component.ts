import { Component } from "@angular/core";
import { MatSidenavModule } from "@angular/material/sidenav";
import { SidenavComponent } from "./sidenav/sidenav.component";
import { SearchPageComponent } from "./search-page/search-page.component";
import { CommonModule } from "@angular/common";
import { MatButtonModule } from "@angular/material/button";
import { MatInputModule } from "@angular/material/input";
import { MatIconModule } from "@angular/material/icon";
import { MatTooltipModule } from "@angular/material/tooltip";
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
  ],
  templateUrl: "./app.component.html",
  styleUrl: "./app.component.css",
})
export class AppComponent {
  title: string = "quDnD";
  name: string = "";
  category: main.Screen = main.Screen.SEARCH;
  navOpened: boolean = true;
  openedPages: Map<main.Screen, main.PageInfo> = new Map();
  currPage: main.PageInfo | SearchPage | main.Screen = { query: "" };

  ngOnInit() {
    document.addEventListener("keydown", (event) => {
      if (event.key === "Escape") {
        this.toggleSidenav();
      }
    });
    this.frontendInit();
  }

  frontendInit() {
    cat.LoadCategories().then((categories) => {});
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
    app.GeneratePage(pageid).then((page) => {
      cat.GetScreen(pageid).then((screen) => {
        this.openedPages.set(screen, page);
        this.currPage = page;
        this.category = screen;
        this.name = page.pageTitle;
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
    }
    this.navOpened = false;
  }

  createCustom() {}
}
