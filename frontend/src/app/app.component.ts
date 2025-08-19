import { Component, inject, ViewChild, ViewChildren } from "@angular/core";
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
import { TabsComponent, TabData } from "./tabs/tabs.component";
import { pageUtils, statblock } from "../../wailsjs/go/models";
import * as cat from "../../wailsjs/go/pageUtils/Categories";
import * as app from "../../wailsjs/go/main/App";
import { filter } from "rxjs";

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
    TabsComponent,
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
  tabs: TabData[] = [];
  isCtrl: boolean = false;
  category: pageUtils.Screen = pageUtils.Screen.SEARCH;
  navOpened: boolean = true;
  openedPages: Map<pageUtils.Screen, statblock.PageInfo> = new Map();
  currPage: statblock.PageInfo | SearchPage | pageUtils.Screen = { query: "" };
  loadingPage: boolean = false;

  ngOnInit() {
    document.addEventListener("keydown", (event) => {
      if (event.key === "Escape") {
        this.toggleSidenav();
      }
    });
    document.addEventListener("keyup", (event) => {
      if (event.ctrlKey) {
        this.isCtrl = false;
      }
    });
    document.addEventListener("keydown", (event) => {
      if (event.ctrlKey) {
        this.isCtrl = true;
      }
    });
    this.frontendInit();
  }

  @ViewChild("tabsComponent") tabsComponent!: TabsComponent;

  ngAfterViewInit() {
    document.addEventListener("keypress", (event) => {
      if (event.key === "T") {
        if (this.isCtrl) {
          event.preventDefault();
          this.tabsComponent.restoreLastTab();
        }
      }
    });
  }

  frontendInit() {
    cat.LoadCategories();
  }

  isSearchPage(): boolean {
    return this.category === "search";
  }

  getScreenName(screen: pageUtils.Screen): string {
    return screen[0].toUpperCase() + screen.slice(1);
  }

  isEmptyPage(): boolean {
    return typeof this.currPage !== "object";
  }

  Number(input: string): number {
    return Number(input);
  }

  goToPage(pageid: number) {
    cat.GetScreen(pageid).then((screen) => {
      this.category = screen;
      this.loadingPage = true;
      this.iframeUrl =
        "https://wiki.cavesofqud.com/Special:Redirect/page/" + pageid;
      this.tabsComponent.selection.set(String(pageid));
      app.GetCachedPage(pageid).then((page) => {
        if (page.exists) {
          this.openedPages.set(screen, page.pageInfo);
          this.name = page.pageInfo.pageTitle;
          this.loadingPage = false;
          this.currPage = page.pageInfo;
          setTimeout(() => {
            let iframe = document.getElementsByClassName(
              "referencePage",
            )[0] as HTMLIFrameElement;
            iframe.src = this.iframeUrl;
            if (
              !this.tabs.filter((val) => {
                return val.id === String(pageid);
              }).length
            ) {
              this.tabs.push({
                id: String(pageid),
                name: this.name,
                icon: String(
                  "https://wiki.cavesofqud.com" + page.pageInfo.imgSrc,
                ),
              });
            }
          }, 100);
        } else {
          app.GeneratePage(pageid).then((page) => {
            this.openedPages.set(screen, page);
            this.currPage = page;
            app.SetCachedPage(page);
            console.log(page.statblock);
            this.name = page.pageTitle;
            this.loadingPage = false;
            setTimeout(() => {
              let iframe = document.getElementsByClassName(
                "referencePage",
              )[0] as HTMLIFrameElement;
              iframe.src = this.iframeUrl;
              if (
                !this.tabs.filter((val) => {
                  return val.id === String(pageid);
                }).length
              ) {
                this.tabs.push({
                  id: String(pageid),
                  name: this.name,
                  icon: String("https://wiki.cavesofqud.com" + page.imgSrc),
                });
              }
            }, 100);
          });
        }
      });
    });
  }

  articleHasDesc(): boolean {
    if (typeof this.currPage === "object")
      return Boolean((this.currPage as statblock.PageInfo).description);
    else return false;
  }

  articleHasImg(): boolean {
    if (typeof this.currPage === "object")
      return Boolean((this.currPage as statblock.PageInfo).imgSrc);
    else return false;
  }

  getDescription(): string {
    if (typeof this.currPage === "object") {
      return (this.currPage as statblock.PageInfo).description ?? "";
    }
    return "";
  }

  getPageInfo(): statblock.PageInfo | undefined {
    if (typeof this.currPage === "object") {
      return this.currPage as statblock.PageInfo;
    }

    return undefined;
  }

  getStatblock(): statblock.Statblock {
    if (typeof this.currPage === "object") {
      let currPage = this.currPage as statblock.PageInfo;
      if (currPage.statblock) {
        return currPage.statblock;
      }
    }

    return new statblock.Statblock();
  }

  getIframeUrl(): SafeResourceUrl {
    return this.sanitizer.bypassSecurityTrustResourceUrl(this.iframeUrl);
  }

  getImgSrc(): string {
    if (typeof this.currPage === "object") {
      return (this.currPage as statblock.PageInfo).imgSrc ?? "";
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
