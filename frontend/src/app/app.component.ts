import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { WorldMapComponent } from "./world-map/world-map.component";
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { HeaderComponent } from "./header/header.component";

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, HeaderComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'frontend';
  // countryMap: Record<string, 'visited' | 'wishlist'> = {
  //   KZ: 'visited',
  //   JP: 'wishlist',
  //   RU: 'visited',
  //   US: 'wishlist',
  //   CN: 'visited',
  // };
  
  // onCountryClick(iso: string) {
  //   console.log("Клик по стране:", iso);
  //   // Можно показать модалку, обновить статус, запросить AI и т.п.
  // }
}
