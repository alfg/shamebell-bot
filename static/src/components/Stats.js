import React, { Component } from 'react';
import StatsEvents from '../StatsEvents';
import './Stats.css';

class Stats extends Component {

  constructor() {
    super();

    this.state = {
      total: 0
    }

    this.stats = new StatsEvents();
    this.stats.on('update', this.onMessage.bind(this));
  }

  onMessage(event) {
    const data = JSON.parse(event.data);
    console.log(data);
    // count = data.total || 0;
    this.setState({ total: data.total });
  }


  render() {
    return(
      <div className="Stats-panel"><span role="img" aria-label="bell">🔔</span>{this.state.total} people shamed!</div>
    )
  }
}

export default Stats;