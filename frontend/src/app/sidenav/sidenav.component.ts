import { Component, input, output, EventEmitter, model } from "@angular/core";
import { CommonModule } from "@angular/common";
import { MatButtonModule } from "@angular/material/button";
import { MatListModule } from "@angular/material/list";
import { main } from "../../../wailsjs/go/models";

@Component({
  selector: "app-sidenav",
  imports: [MatButtonModule, CommonModule, MatListModule],
  templateUrl: "./sidenav.component.html",
  styleUrl: "./sidenav.component.css",
})
export class SidenavComponent {
  categories = [
    main.Screen.SEARCH,
    main.Screen.CHARACTER,
    main.Screen.CONCEPTS,
    main.Screen.CREATURES,
    main.Screen.CUSTOM,
    main.Screen.ITEMS,
    main.Screen.LIQUIDS,
    main.Screen.LORE,
    main.Screen.MECHANICS,
    main.Screen.MUTATIONS,
    main.Screen.OTHER,
  ];
  change = output<void>();
  selectedScreen = model(main.Screen.SEARCH);

  getScreenName(screen: main.Screen): string {
    return screen[0].toUpperCase() + screen.slice(1);
  }

  ngOnInit() {
    this.change.emit();
  }

  selectScreen(screen: main.Screen) {
    this.selectedScreen.set(screen);
    this.change.emit();
  }
}
