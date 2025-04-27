import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';
import { environment } from '../environments/environment';
@Injectable({
  providedIn: 'root'
})
export class StatisticsService {
  private apiUrl = environment.apiUrl;
  constructor(private http: HttpClient) { }

  getTopVisitedCountries(): Promise<any> {
    return firstValueFrom(this.http.get<any>(`${this.apiUrl}/statistics/top-visited`, { withCredentials: true }));
  }

  getTopWishlistCountries(): Promise<any> {
    return firstValueFrom(this.http.get<any>(`${this.apiUrl}/statistics/top-wish-list`, { withCredentials: true }));
  }
}
