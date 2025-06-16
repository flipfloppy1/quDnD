import { Component, input, signal } from "@angular/core";
import { KeyValuePipe, CommonModule } from "@angular/common";
import { MatExpansionModule } from "@angular/material/expansion";
import { MatTooltipModule } from "@angular/material/tooltip";
import { statblock } from "../../../wailsjs/go/models";

@Component({
  selector: "app-statblock",
  imports: [MatExpansionModule, KeyValuePipe, CommonModule, MatTooltipModule],
  templateUrl: "./statblock.component.html",
  styleUrl: "./statblock.component.css",
})
export class StatblockComponent {
  statblock = input(new statblock.Statblock());
  itemClick = signal<number>;

  hasDetailedStats(): boolean {
    return (
      Object.entries(this.statblock().stats).filter((val) => {
        return (
          val[0] != "hp" &&
          val[0] != "ac" &&
          val[0] != "speed" &&
          val[0] != "cha" &&
          val[0] != "con" &&
          val[0] != "int" &&
          val[0] != "wis" &&
          val[0] != "dex" &&
          val[0] != "str"
        );
      }).length > 0
    );
  }

  detailedStats(): Map<string, string> {
    return new Map(
      Object.entries(this.statblock().stats).filter((val) => {
        return (
          val[0] != "hp" &&
          val[0] != "ac" &&
          val[0] != "speed" &&
          val[0] != "cha" &&
          val[0] != "con" &&
          val[0] != "int" &&
          val[0] != "wis" &&
          val[0] != "dex" &&
          val[0] != "str"
        );
      }),
    );
  }

  hasAbilityScores(): boolean {
    return (
      Object.entries(this.statblock().stats).filter((val) => {
        return (
          val[0] == "cha" ||
          val[0] == "con" ||
          val[0] == "int" ||
          val[0] == "wis" ||
          val[0] == "dex" ||
          val[0] == "str"
        );
      }).length == 6
    );
  }

  abilityScores(): Map<string, string> {
    return new Map(
      Object.entries(this.statblock().stats).filter((val) => {
        return (
          val[0] == "cha" ||
          val[0] == "con" ||
          val[0] == "int" ||
          val[0] == "wis" ||
          val[0] == "dex" ||
          val[0] == "str"
        );
      }),
    );
  }

  getAbilityMod(ability: string): string {
    let mod: number = Math.ceil(
      (Number(this.statblock().stats[ability]) - 10) / 2,
    );
    if (mod < 0) {
      return String(mod);
    } else {
      return "+" + mod;
    }
  }
}
