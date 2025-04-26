import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { MatSidenavModule } from "@angular/material/sidenav";
import { SidenavComponent } from "./sidenav/sidenav.component";
import { CommonModule } from "@angular/common";
import { MatButtonModule } from "@angular/material/button";
import { MatIconModule } from "@angular/material/icon";
import { MatTooltipModule } from "@angular/material/tooltip";
import { main } from "../../wailsjs/go/models";

@Component({
  selector: "app-root",
  imports: [
    RouterOutlet,
    MatSidenavModule,
    SidenavComponent,
    CommonModule,
    MatButtonModule,
    MatIconModule,
  ],
  templateUrl: "./app.component.html",
  styleUrl: "./app.component.css",
})
export class AppComponent {
  title: string = "quDnD";
  name: string = "";
  category: main.Screen = main.Screen.OTHER;
  navOpened: boolean = true;

  ngOnInit() {}

  toggleSidenav() {
    this.navOpened = !this.navOpened;
  }

  createCustom() {}
}
