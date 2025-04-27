import { AfterViewInit, Component, ElementRef, EventEmitter, Input, OnDestroy, Output, ViewChild } from '@angular/core';
import * as am5 from '@amcharts/amcharts5';
import * as am5map from '@amcharts/amcharts5/map';
import am5geodata_world from '@amcharts/amcharts5-geodata/worldLow';
import { CountryStatusService } from '../country-status.service';
import { AuthService } from '../auth.service';
import { AiService } from '../ai.service';
import { CommonModule } from '@angular/common';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { marked } from 'marked';

@Component({
  selector: 'app-world-map',
  imports: [CommonModule],
  templateUrl: './world-map.component.html',
  styleUrls: ['./world-map.component.css']
})
export class WorldMapComponent implements AfterViewInit, OnDestroy {
  @ViewChild('chartdiv', { static: true }) chartdiv!: ElementRef<HTMLDivElement>;
  @Input() countriesStatus: Record<string, 'visited' | 'wishlist'> = {};
  @Output() countryClicked = new EventEmitter<string>();

  private root!: am5.Root;
  private polygonSeries!: am5map.MapPolygonSeries;
  selectedCountryISO: string | null = null;
  selectedCountryName: string | null = null;
  countryInfo: SafeHtml = '';
  isAuthenticated: boolean = false;

  constructor(
    private countryStatusService: CountryStatusService,
    private authService: AuthService,
    private aiService: AiService,
    private sanitizer: DomSanitizer
  ) {}

  ngAfterViewInit(): void {
    this.authService.profile().subscribe({
      next: () => {
        this.isAuthenticated = true;
        this.initMap();
        this.loadCountryStatuses();
      },
      error: () => {
        this.isAuthenticated = false;
        this.initMap();
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
      const dataContext = ev.target.dataItem?.dataContext as any;
      const id = dataContext?.id;
      const name = dataContext?.name;

      if (id && this.isAuthenticated) {
        this.selectedCountryISO = id;
        this.selectedCountryName = name;
        this.countryClicked.emit(id);
        this.countryInfo = '';
      }
    });
  }

  private loadCountryStatuses(): void {
    if (!this.isAuthenticated) return;

    this.countryStatusService.getCountryStatuses().subscribe({
      next: (statuses) => {
        this.countriesStatus = statuses;
        if (this.polygonSeries) {
           this.polygonSeries.data.setAll(this.polygonSeries.data.values);
        }
      },
      error: (err) => console.error('Failed to load country statuses:', err)
    });
  }

  private async fetchCountryInfo(countryName: string): Promise<void> {
    this.aiService.ask(`Расскажи о стране ${countryName} 500 символов макс и дай ссылку на википедию. Пиши сразу не надо вот вам и т д`).subscribe({
      next: async (response) => {
        try {
          const htmlContent = await marked(response.answer);
          console.log('Полученный HTML:', htmlContent);
          this.countryInfo = this.sanitizer.bypassSecurityTrustHtml(htmlContent);
        } catch (error) {
          console.error('Ошибка при обработке Markdown:', error);
          this.countryInfo = this.sanitizer.bypassSecurityTrustHtml('<p>Не удалось получить информацию о стране из-за ошибки форматирования.</p>');
        }
      },
      error: (err) => {
        console.error('Ошибка при запросе информации о стране:', err);
        this.countryInfo = this.sanitizer.bypassSecurityTrustHtml('<p>Не удалось получить информацию о стране.</p>');
      }
    });
  }

  onAskCountryInfo(): void {
    if (this.selectedCountryName) {
      this.fetchCountryInfo(this.selectedCountryName);
    }
  }

  setStatus(status: 'visited' | 'wishlist' | 'none'): void {

    if (this.selectedCountryISO && this.isAuthenticated) {
      this.updateCountryStatus(this.selectedCountryISO, status);
      this.selectedCountryISO = null;
      this.selectedCountryName = null;
      this.countryInfo = '';
    }
  }

  private updateCountryStatus(countryISO: string, status: 'visited' | 'wishlist' | 'none'): void {
    this.countryStatusService.setCountryStatus(countryISO, status).subscribe({
      next: () => this.loadCountryStatuses(),
      error: (err) => console.error('Failed to update country status:', err)
    });
  }

  closeModal(): void {
    this.selectedCountryISO = null; // Скрыть модальное окно
    this.selectedCountryName = null;
    this.countryInfo = '';
  }

  ngOnDestroy(): void {
    this.root?.dispose();
  }
}