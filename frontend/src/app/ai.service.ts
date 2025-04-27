import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../environments/environment';
@Injectable({
  providedIn: 'root',
})
export class AiService {

  private apiUrl = environment.apiUrl;

  constructor(private http: HttpClient) {}

  ask(prompt: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/ai/ask?prompt=${encodeURIComponent(prompt)}`, { withCredentials: true });
  }
}

