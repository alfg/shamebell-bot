import React, { Component } from 'react';
import StatsEvents from '../StatsEvents';
import './Stats.css';

class Stats extends Component {

  constructor() {
    super();

    this.state = {
      total: 0,
      servers: 0,
    }

    this.stats = new StatsEvents();
    this.stats.on('update', this.onMessage.bind(this));
  }

  onMessage(event) {
    const data = JSON.parse(event.data);
    this.setState({ total: data.total || 0, servers: data.guilds || 0 });
  }


  render() {
    return(
      <div className="Stats-panel">
        <span role="img" aria-label="bell">🔔</span>{this.state.total} people shamed on {this.state.servers} servers!
      </div>
    )
  }
}

export default Stats;