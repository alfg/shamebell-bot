import React, { Component } from 'react';
import profile from '../img/profile.png';
import './App.css';
import Constants from '../Constants';
import Stats from './Stats';
import Footer from './Footer';

class App extends Component {

  addToDiscordHandler(e) {
    window.open(Constants.AUTH_URL, '', 'height=800, width=500');
  }

  playAudio = () => {
    this.refs.audio.play();
  }

  render() {
    return (
      <div className="App">
        <div className="App-content">
          <div className="App-intro">
            <img src={profile} className="App-logo" alt="logo" onClick={this.playAudio} />
            <h1>!shamebell</h1>
            Shame your friends on <a href={Constants.DISCORD_URL}>Discord</a> for doing stupid things
          </div>

          <a className="App-button" onClick={this.addToDiscordHandler}>
            Add to Discord
          </a>
          <p className="App-instruction">Just type <span className="command">!shame</span> while in a voice channel!</p>

          <Stats />
          <audio preload src={Constants.AUDIO_FILE} type="audio/mp3" ref="audio" />
        </div>

        <Footer />
      </div>
    );
  }
}

export default App;