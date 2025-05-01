import { Component, ViewEncapsulation, output } from "@angular/core";
import { CommonModule } from "@angular/common";
import { FormsModule } from "@angular/forms";
import { MatButtonModule } from "@angular/material/button";
import { MatInputModule } from "@angular/material/input";
import { MatIconModule } from "@angular/material/icon";
import * as app from "../../../wailsjs/go/main/App";
import { main } from "../../../wailsjs/go/models";

@Component({
  selector: "app-search-page",
  imports: [
    MatButtonModule,
    MatInputModule,
    MatIconModule,
    CommonModule,
    FormsModule,
  ],
  encapsulation: ViewEncapsulation.ShadowDom,
  templateUrl: "./search-page.component.html",
  styleUrl: "./search-page.component.css",
})
export class SearchPageComponent {
  searchText: string = "";
  results: main.RestPageSearchResults[] = [];
  changePage = output<number>({ alias: "search" });

  changeToPage(pageid: number) {
    this.changePage.emit(pageid);
  }

  submit() {
    app.SearchForPage(this.searchText).then((result) => {
      this.results = result.pages;
    });
  }
}
