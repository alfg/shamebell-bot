import { EventEmitter } from 'events';

class Stats extends EventEmitter {
    constructor() {
        super();

        if (EventSource) {
            let eventSource = new EventSource('/events');
            eventSource.onmessage = this.receivedMessage.bind(this);
        }

    }

    receivedMessage(event) {
        console.log(event);
        this.emit('update');
    }

}

export default Stats;

