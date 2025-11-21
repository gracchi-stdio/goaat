const formGroupStyles = new CSSStyleSheet();

formGroupStyles.replaceSync(`

  
  .error-message {
    color: var(--error-message-color, red);
  }

  .error-message:empty {
    display: none;
  }

`);

class FormGroup extends HTMLElement {
  #input;

  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.adoptedStyleSheets = [formGroupStyles];
    this.shadowRoot.innerHTML = `
      <slot></slot>
      <span class="error-message" 
            aria-live="polite"
            part="error-message">
      </span>
    `
  }

  #handleInvalid(e) {
    e.preventDefault();
  }

  #handleInput() {
    if(this.#input) {
      console.log('hello')
      this.#errorMessage.textContent = '';
    }
  }

  #handleBlur() {
    // if input is NOT valid, show an error message
    if(this.#input && !this.#input.validity.valid) {
      this.#errorMessage.textContent = this.#customErrorMessage[this.#getFirstInvalid(this.#input.validity)];
    }
  }

  connectedCallback() {
    this.#input = this.querySelector("input, textarea");

    if(this.#input) {
      // turn off the browser validation popup
      this.#input.addEventListener('invalid', this.#handleInvalid.bind(this))
      
      // hide error messages while user is typing
      this.#input.addEventListener('input', this.#handleInput.bind(this))
      
      // validate and show error messages when a user leaves a field
      this.#input.addEventListener('blur', this.#handleBlur.bind(this))
    }
  }

  disconnectedCallback() {
    // clean up all event listeners
    if(this.#input) {
      this.#input.removeEventListener('invalid',  this.#handleInvalid)
      this.#input.removeEventListener('input', this.#handleInput)
      this.#input.removeEventListener('blur', this.#handleBlur)
    }
  }

  get #errorMessage() {
    return this.shadowRoot.querySelector('.error-message');
  }

  #getFirstInvalid(validityState) {
    for(const key in validityState) {
      if(validityState[key]) {
        return key;
      }
    }
  }

  get #customErrorMessage() {
    return {
      valueMissing: this.getAttribute('value-missing-message') || 'This field is required',
      tooLong: this.getAttribute('too-long-message') || 'This field is too long',
      tooShort: this.getAttribute('too-short-message') || 'This field is too short',
      rangeOverflow: this.getAttribute('range-overflow-message') || 'This field has a number that is too big',
      rangeUnderflow: this.getAttribute('range-underflow-message') || 'This field has a number that is too small',
      typeMismatch: this.getAttribute('type-mismatch-message') || 'This field is the wrong type',
      patternMismatch: this.getAttribute('pattern-mismatch-message') || 'This fields value does not match the pattern',
    }
  }
}  




customElements.define("form-group", FormGroup);