import {
  AfterViewInit, Component, ElementRef, EventEmitter, Input, OnDestroy, Output, ViewChild
} from '@angular/core';
import * as am5 from "@amcharts/amcharts5";
import * as am5map from "@amcharts/amcharts5/map";
import am5geodata_world from "@amcharts/amcharts5-geodata/worldLow";

@Component({
  selector: 'app-world-map',
  imports: [],
  templateUrl: './world-map.component.html',
  styleUrl: './world-map.component.css'
})
export class WorldMapComponent implements AfterViewInit, OnDestroy {
  @ViewChild('chartdiv', { static: true }) chartdiv!: ElementRef<HTMLDivElement>;
  @Input() countriesStatus: Record<string, 'visited' | 'wishlist'> = {};
  @Output() countryClicked = new EventEmitter<string>();

  private root!: am5.Root;

  ngAfterViewInit(): void {
    this.root = am5.Root.new(this.chartdiv.nativeElement);

    const chart = this.root.container.children.push(
      am5map.MapChart.new(this.root, {
        panX: "none",
        panY: "none",
        projection: am5map.geoEqualEarth(),
      })
    );

    const polygonSeries = chart.series.push(
      am5map.MapPolygonSeries.new(this.root, {
        geoJSON: am5geodata_world
      })
    );

    polygonSeries.mapPolygons.template.setAll({
      tooltipText: "{name}",
      interactive: true
    });

    polygonSeries.mapPolygons.template.states.create("hover", {
      fill: am5.color(0x677935)
    });

    // Цвета по статусу
    polygonSeries.mapPolygons.template.adapters.add("fill", (fill, target) => {
      const id = (target.dataItem?.dataContext as any).id;

      if (id && this.countriesStatus[id]) {
        return this.countriesStatus[id] === "visited"
          ? am5.color(0x4caf50) // зелёный
          : am5.color(0xffc107); // жёлтый
      }
      return fill;
    });

    // Клик по стране
    polygonSeries.mapPolygons.template.events.on("click", (ev) => {
      const id = (ev.target.dataItem?.dataContext as any).id;
      if (id) {
        this.countryClicked.emit(id);
      }
    });
  }

  ngOnDestroy(): void {
    this.root?.dispose();
  }
}