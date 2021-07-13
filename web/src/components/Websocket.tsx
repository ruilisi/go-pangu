import React from 'react';
import PropTypes from 'prop-types';

type PropType = {
  url: string;
  onClose: ((ev: CloseEvent) => any) | null;
  onError: ((ev: Event) => any) | null;
  onMessage: ((ev: MessageEvent) => any) | null;
  onOpen: ((ev: Event) => any) | null;
  reconnect: boolean;
  protocol: string;
  reconnectIntervalInMilliSeconds: number;
};

type StateType = {
  ws: WebSocket;
  attempts: number;
};

class Websocket extends React.Component<PropType, StateType> {
  static defaultProps: Partial<PropType> = {
    reconnect: true,
  };

  static propTypes = {
    url: PropTypes.string.isRequired,
    onMessage: PropTypes.func.isRequired,
    onOpen: PropTypes.func,
    onClose: PropTypes.func,
    onError: PropTypes.func,
    debug: PropTypes.bool,
    reconnect: PropTypes.bool,
    protocol: PropTypes.string,
    reconnectIntervalInMilliSeconds: PropTypes.number,
  };

  attempts: number = 1;
  ws: WebSocket | undefined;
  timeoutID:any = setTimeout(() => {});

  initWebsocket(attempts: number) {
    this.ws = new window.WebSocket(this.props.url, this.props.protocol);
    this.attempts = attempts;
    this.setupWebsocket();
  }

  generateInterval = (k: number) => {
    if (this.props.reconnectIntervalInMilliSeconds > 0) {
      return this.props.reconnectIntervalInMilliSeconds;
    }
    return Math.min(30, 2 ** k - 1) * 1000;
  };

  setupWebsocket = () => {
    const websocket = this.ws!;

    websocket.onopen = (e) => {
      if (typeof this.props.onOpen === 'function') this.props.onOpen(e);
    };

    websocket.onerror = (e) => {
      if (typeof this.props.onError === 'function') this.props.onError(e);
    };

    websocket.onmessage = (e) => {
      if (typeof this.props.onMessage === 'function') this.props.onMessage(e.data);
    };

    websocket.onclose = (e) => {
      if (typeof this.props.onClose === 'function') this.props.onClose(e);
      if (this.props.reconnect) {
        const time = this.generateInterval(this.attempts);
        this.timeoutID = setTimeout(() => {
          this.initWebsocket(this.attempts + 1);
        }, time);
      }
    };
  };

  componentDidUpdate(prevProps: PropType) {
    if (this.props.url !== prevProps.url || this.props.protocol !== prevProps.protocol) {
      this.close();
      this.initWebsocket(1);
    }
  }

  componentDidMount() {
    this.initWebsocket(1);
  }

  componentWillUnmount() {
    this.close();
  }

  close() {
    clearTimeout(this.timeoutID);
    const websocket = this.ws;
    if (!websocket) return;

    websocket.onclose = (e) => {
      if (typeof this.props.onClose === 'function') this.props.onClose(e);
      // removed reconnect logic
    };
    websocket.close();
  }

  sendMessage = (message: string) => {
    this.ws?.send(message);
  };

  render() {
    return null;
  }
}

export default Websocket;
