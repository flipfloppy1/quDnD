import {
  Component,
  ElementRef,
  ViewChild,
  ViewEncapsulation,
  inject,
  output,
} from "@angular/core";
import { CommonModule } from "@angular/common";
import { FormsModule } from "@angular/forms";
import { MatButtonModule } from "@angular/material/button";
import { MatInputModule, MatInput } from "@angular/material/input";
import { MatIconModule } from "@angular/material/icon";
import { MatProgressSpinnerModule } from "@angular/material/progress-spinner";
import {
  MatAutocomplete,
  MatAutocompleteModule,
  MatAutocompleteTrigger,
} from "@angular/material/autocomplete";
import * as app from "../../../wailsjs/go/main/App";
import * as cat from "../../../wailsjs/go/pageUtils/Categories";
import { pageUtils } from "../../../wailsjs/go/models";

@Component({
  selector: "app-search-page",
  imports: [
    MatButtonModule,
    MatInputModule,
    MatIconModule,
    MatAutocompleteModule,
    CommonModule,
    FormsModule,
    MatProgressSpinnerModule,
  ],
  //encapsulation: ViewEncapsulation.ShadowDom,
  templateUrl: "./search-page.component.html",
  styleUrl: "./search-page.component.css",
  standalone: true,
})
export class SearchPageComponent {
  @ViewChild(MatAutocompleteTrigger) auto!: MatAutocompleteTrigger;
  searching: boolean = false;
  noResults: boolean = false;
  searchText: string = "";
  prevSearchText: string = "";
  typeahead: pageUtils.PageData[] = [];
  results: pageUtils.RestPageSearchResults[] = [];
  changePage = output<number>({ alias: "search" });

  changeToPage(pageid: number) {
    this.changePage.emit(pageid);
  }

  fuzzySearch() {
    if (this.searchText != this.prevSearchText && this.searchText.length > 2) {
      cat.FuzzySearch(this.searchText).then((result) => {
        this.typeahead = result;
      });
      this.prevSearchText = this.searchText;
    }
  }

  submit() {
    this.auto.closePanel();
    this.searching = true;
    app.SearchForPage(this.searchText).then((result) => {
      this.results = result.pages;
      this.noResults = !this.results.length;
      this.searching = false;
    });
  }
}
