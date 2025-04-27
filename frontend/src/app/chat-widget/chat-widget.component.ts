import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { DomSanitizer } from '@angular/platform-browser';
import { AiService } from '../ai.service';
import { AuthService } from '../auth.service';
@Component({
  selector: 'app-chat-widget',
  imports: [CommonModule, FormsModule],
  templateUrl: './chat-widget.component.html',
  styleUrls: ['./chat-widget.component.css']
})
export class ChatWidgetComponent {
  isChatOpen = false;
  isLoggedIn: boolean = false;
  userInput = '';
  messages: {role: string, text: string}[] = [];

  constructor(private aiService: AiService, private sanitizer: DomSanitizer, private authService: AuthService) {}

  toggleChat() {
    this.isChatOpen = !this.isChatOpen;
  }

  ngOnInit(): void {
    this.authService.checkLoginStatus();
    
    this.authService.isLoggedIn$.subscribe(isLoggedIn => {
      this.isLoggedIn = isLoggedIn;
    });
  }
  

  sendMessage() {
    if (!this.userInput.trim()) return;
  
    const userMessage = this.userInput;
    this.messages.push({ role: 'Пользователь', text: userMessage });
  
    this.userInput = '';
  
    this.aiService.ask(`Ты - дружелюбный и знающий travel-гид. Всегда отвечай в этом стиле. Твоя цель - вдохновлять людей путешествовать и делиться интересной и полезной информацией о различных местах, культурах и аспектах путешествий. Общайся так, будто ты лично консультируешь путешественника или рассказываешь увлекательную историю о месте.${userMessage}`)
      .subscribe({
        next: async (response) => {
          try {
            const answer = response.answer;
            this.messages.push({ role: 'Бот', text: answer });
          } catch (error) {
            console.error('Ошибка обработки ответа:', error);
            this.messages.push({ role: 'Бот', text: 'Ошибка обработки ответа от ИИ.' });
          }
        },
        error: (err) => {
          console.error('Ошибка запроса к ИИ сервису:', err);
          this.messages.push({ role: 'Бот', text: 'Ошибка соединения с ИИ.' });
        }
      });
  }
}
