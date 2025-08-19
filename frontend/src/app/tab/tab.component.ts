import { Component, input, output } from "@angular/core";
import { MatIconModule } from "@angular/material/icon";
import { MatButtonModule } from "@angular/material/button";
import { CommonModule } from "@angular/common";

@Component({
  selector: "app-tab",
  imports: [CommonModule, MatButtonModule, MatIconModule],
  templateUrl: "./tab.component.html",
  styleUrl: "./tab.component.css",
})
export class TabComponent {
  selected = input(false);
  name = input("");
  id = input("");
  icon = input("");
  close = output();
  selectTab = output();

  closeTab(event: Event) {
    event.preventDefault();
    this.close.emit();
  }
}
