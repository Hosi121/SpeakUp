import api from './api'; // Importing the existing axios instance

interface WebSocketMessage {
  type: string;
  content: any;
}

class WebSocketService {
  public socket: WebSocket | null = null;
  private readonly baseUrl: string;

  constructor(baseUrl?: string) {
    if (baseUrl) {
      this.baseUrl = baseUrl.replace(/^http/, 'ws');
    } else if (api.defaults.baseURL) {
      this.baseUrl = api.defaults.baseURL.replace(/^http/, 'ws');
    } else {
      throw new Error("Base URL is not provided and not defined in the API configuration");
    }
  }

  public connect(): void {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No token found");
      return;
    }

    this.socket = new WebSocket(`${this.baseUrl}/ws?token=${token}`);

    this.socket.onopen = this.onOpen.bind(this);
    this.socket.onmessage = this.onMessage.bind(this);
    this.socket.onerror = this.onError.bind(this);
    this.socket.onclose = this.onClose.bind(this);
  }

  private onOpen(event: Event): void {
    console.log("WebSocket connection established", event);
  }

  private onMessage(event: MessageEvent): void {
    console.log("Received message:", event.data);
    // Handle incoming messages here
  }

  private onError(event: Event): void {
    console.error("WebSocket error:", event);
  }

  private onClose(event: CloseEvent): void {
    console.log("WebSocket connection closed", event);
    this.socket = null;
  }

  public sendMessage(message: WebSocketMessage): void {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify(message));
    } else {
      console.error("WebSocket is not connected");
    }
  }

  public disconnect(): void {
    if (this.socket) {
      this.socket.close();
      this.socket = null;
    }
  }

  public isConnected(): boolean {
    return this.socket !== null && this.socket.readyState === WebSocket.OPEN;
  }
}

// シングルトンインスタンスとしてエクスポートする場合は、デフォルトのbaseUrlを使用
export default new WebSocketService();

// クラスをエクスポートして、個別にインスタンス化できるようにする
export { WebSocketService };
