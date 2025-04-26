import { AfterViewInit, Component, ElementRef, EventEmitter, Input, OnDestroy, Output, ViewChild } from '@angular/core';
import * as am5 from '@amcharts/amcharts5';
import * as am5map from '@amcharts/amcharts5/map';
import am5geodata_world from '@amcharts/amcharts5-geodata/worldLow';
import { CountryStatusService } from '../country-status.service';
import { AuthService } from '../auth.service'; // Импортируйте AuthService
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-world-map',
  imports: [CommonModule],
  templateUrl: './world-map.component.html',
  styleUrl: './world-map.component.css'
})
export class WorldMapComponent implements AfterViewInit, OnDestroy {
  @ViewChild('chartdiv', { static: true }) chartdiv!: ElementRef<HTMLDivElement>;
  @Input() countriesStatus: Record<string, 'visited' | 'wishlist'> = {};
  @Output() countryClicked = new EventEmitter<string>();

  private root!: am5.Root;
  private polygonSeries!: am5map.MapPolygonSeries;
  selectedCountry: string | null = null;
  isAuthenticated: boolean = false; // Для отслеживания статуса аутентификации

  constructor(
    private countryStatusService: CountryStatusService,
    private authService: AuthService
  ) {}

  ngAfterViewInit(): void {
    // Проверка аутентификации
    this.authService.profile().subscribe({
      next: () => {
        this.isAuthenticated = true;
        this.initMap();
        this.loadCountryStatuses();
      },
      error: () => {
        this.isAuthenticated = false;
        this.initMap(); // Карта может быть показана, но без данных
      }
    });
  }

  private initMap(): void {
    this.root = am5.Root.new(this.chartdiv.nativeElement);

    const chart = this.root.container.children.push(
      am5map.MapChart.new(this.root, {
        panX: 'none',
        panY: 'none',
        projection: am5map.geoEqualEarth(),
      })
    );

    this.polygonSeries = chart.series.push(
      am5map.MapPolygonSeries.new(this.root, {
        geoJSON: am5geodata_world
      })
    );

    this.polygonSeries.mapPolygons.template.setAll({
      tooltipText: '{name}',
      interactive: true
    });

    this.polygonSeries.mapPolygons.template.states.create('hover', {
      fill: am5.color(0x677935)
    });

    this.polygonSeries.mapPolygons.template.adapters.add('fill', (fill, target) => {
      const id = (target.dataItem?.dataContext as any).id;
      if (id && this.countriesStatus[id]) {
        return this.countriesStatus[id] === 'visited'
          ? am5.color(0x4caf50)
          : am5.color(0xffc107);
      }
      return fill;
    });

    this.polygonSeries.mapPolygons.template.events.on('click', (ev) => {
      const id = (ev.target.dataItem?.dataContext as any).id;
      if (id && this.isAuthenticated) {
        this.selectedCountry = id;
        this.countryClicked.emit(id);
      }
    });
  }

  private loadCountryStatuses(): void {
    if (!this.isAuthenticated) return;

    this.countryStatusService.getCountryStatuses().subscribe({
      next: (statuses) => {
        this.countriesStatus = statuses;
        this.polygonSeries?.data.setAll(this.polygonSeries.data.values);
      },
      error: (err) => console.error('Failed to load country statuses:', err)
    });
  }

  setStatus(status: 'visited' | 'wishlist' | 'none'): void {
    if (this.selectedCountry && this.isAuthenticated) {
      this.updateCountryStatus(this.selectedCountry, status);
      this.selectedCountry = null;
    }
  }

  private updateCountryStatus(countryISO: string, status: 'visited' | 'wishlist' | 'none'): void {
    this.countryStatusService.setCountryStatus(countryISO, status).subscribe({
      next: () => this.loadCountryStatuses(),
      error: (err) => console.error('Failed to update country status:', err)
    });
  }

  ngOnDestroy(): void {
    this.root?.dispose();
  }
}