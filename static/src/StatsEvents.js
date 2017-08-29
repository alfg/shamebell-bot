import { EventEmitter } from 'events';

class StatsEvents extends EventEmitter {
    constructor() {
        super();

        if (EventSource) {
            let eventSource = new EventSource('/events');
            eventSource.onmessage = this.receivedMessage.bind(this);
        }

    }

    receivedMessage(event) {
      this.emit('update', event);
    }

}

export default StatsEvents;

