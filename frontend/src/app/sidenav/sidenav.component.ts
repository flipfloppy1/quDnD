import { Component, input, output, EventEmitter, model } from "@angular/core";
import { CommonModule } from "@angular/common";
import { MatButtonModule } from "@angular/material/button";
import { MatListModule } from "@angular/material/list";
import { pageUtils } from "../../../wailsjs/go/models";

@Component({
  selector: "app-sidenav",
  imports: [MatButtonModule, CommonModule, MatListModule],
  templateUrl: "./sidenav.component.html",
  styleUrl: "./sidenav.component.css",
})
export class SidenavComponent {
  categories = [
    pageUtils.Screen.SEARCH,
    pageUtils.Screen.CHARACTER,
    pageUtils.Screen.CONCEPTS,
    pageUtils.Screen.CREATURES,
    pageUtils.Screen.CUSTOM,
    pageUtils.Screen.ITEMS,
    pageUtils.Screen.LIQUIDS,
    pageUtils.Screen.LORE,
    pageUtils.Screen.MECHANICS,
    pageUtils.Screen.MUTATIONS,
    pageUtils.Screen.OTHER,
  ];
  change = output<void>();
  selectedScreen = model(pageUtils.Screen.SEARCH);

  getScreenName(screen: pageUtils.Screen): string {
    return screen[0].toUpperCase() + screen.slice(1);
  }

  ngOnInit() {
    this.change.emit();
  }

  selectScreen(screen: pageUtils.Screen) {
    this.selectedScreen.set(screen);
    this.change.emit();
  }
}
