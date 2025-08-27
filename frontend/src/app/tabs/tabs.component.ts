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

  navLeft() {
    let idx;
    const tabs = this.tabs();

    if (!tabs.length) {
      return;
    }

    for (let i = 0; i < tabs.length; i++) {
      if (this.selection() === tabs[i].id) {
        idx = i;
      }
    }

    if (idx == undefined) return;

    if (idx - 1 >= 0) {
      this.selection.set(tabs[idx - 1].id);
    } else {
      this.selection.set(tabs[tabs.length - 1].id);
    }
  }

  navRight() {
    let idx;
    const tabs = this.tabs();

    if (!tabs.length) {
      return;
    }

    for (let i = 0; i < tabs.length; i++) {
      if (this.selection() === tabs[i].id) {
        idx = i;
      }
    }

    if (idx == undefined) return;

    if (idx + 1 < tabs.length) {
      this.selection.set(tabs[idx + 1].id);
    } else {
      this.selection.set(tabs[0].id);
    }
  }

  restoreLastTab() {
    const prevTab = this.prevTabs.pop();
    if (prevTab) {
      this.tabs().push(prevTab);
      this.selection.set(prevTab.id);
    }
  }

  closeTab(id: string) {
    let tabs = this.tabs();
    let elemList;
    let elemIdx;
    for (let i = 0; i < tabs.length; i++) {
      if (tabs[i].id === id) {
        elemList = tabs.splice(i, 1);
        elemIdx = i;
      }
    }
    if (!elemList || !elemList.length) {
      return;
    }
    this.prevTabs.push(elemList[0]);
    if (tabs.length === 0) {
      this.selection.set("");
      return;
    }

    this.tabs.set(tabs);

    if ((elemIdx as number) >= tabs.length) {
      (elemIdx as number)--;
    }
    this.selection.set(tabs[elemIdx as number].id);
  }
}
