import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class AiService {

  private apiUrl = 'http://localhost:1488/api/ai/ask';

  constructor(private http: HttpClient) {}

  ask(prompt: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}?prompt=${encodeURIComponent(prompt)}`, { withCredentials: true });
  }
}

