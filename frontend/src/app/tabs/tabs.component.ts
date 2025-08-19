import { Component, input, model } from "@angular/core";
import { CommonModule } from "@angular/common";
import { TabComponent } from "../tab/tab.component";
import {
  CdkDragDrop,
  CdkDropList,
  CdkDrag,
  moveItemInArray,
} from "@angular/cdk/drag-drop";

export interface TabData {
  name: string;
  id: string;
  icon: string;
}

@Component({
  selector: "app-tabs",
  imports: [CdkDropList, CdkDrag, TabComponent],
  templateUrl: "./tabs.component.html",
  styleUrl: "./tabs.component.css",
})
export class TabsComponent {
  tabs = model<TabData[]>([]);
  prevTabs: TabData[] = [];
  selection = model("");

  drop(event: CdkDragDrop<TabData[]>) {
    moveItemInArray(this.tabs(), event.previousIndex, event.currentIndex);
  }

  restoreLastTab() {
    const prevTab = this.prevTabs.pop();
    if (prevTab) {
      this.tabs().push(prevTab);
    }
  }

  closeTab(id: string) {
    this.prevTabs.push(
      this.tabs().filter((val) => {
        return val.id === id;
      })[0],
    );
    this.tabs.set(
      this.tabs().filter((val) => {
        return val.id !== id;
      }),
    );
  }
}
