function drawLine(x, y, width, height) {
  const line = document.createElement('div');
  line.style.position = 'absolute';
  line.style.left = x;
  line.style.top = y;
  line.style.width = width;
  line.style.height = height;
  line.style.background = 'black';
  document.body.appendChild(line);
}

window.onload = function () {
  const thickness = '0.2mm';       // Thickness of lines
  const length = '5mm';            // Length of crop marks
  const pixel = '0.26mm';          // Approx. 1px correction (96 DPI)

  // --- CORNERS ---

  // Top-left corner
  drawLine('calc(1cm - ' + length + ' - ' + pixel + ')', 'calc(1cm - ' + pixel + ')', length, thickness); // horizontal
  drawLine('calc(1cm - ' + pixel + ')', 'calc(1cm - ' + length + ' - ' + pixel + ')', thickness, length); // vertical

  // Top-right corner
  drawLine('calc(16cm + ' + pixel + ')', 'calc(1cm - ' + pixel + ')', length, thickness); // horizontal (fixed)
  drawLine('calc(16cm - ' + thickness + ' + ' + pixel + ')', 'calc(1cm - ' + length + ' - ' + pixel + ')', thickness, length); // vertical (fixed)

  // Bottom-left corner
  drawLine('calc(1cm - ' + length + ' - 1 * ' + pixel + ')', 'calc(12cm - 0 * ' + pixel + ')', length, thickness); // horizontal (double left shift)
  drawLine('calc(1cm - ' + pixel + ')', 'calc(12cm + ' + pixel + ')', thickness, length); // vertical

  // Bottom-right corner
  drawLine('calc(16cm + ' + pixel + ')', 'calc(12cm)', length, thickness); // horizontal (fixed)
  drawLine('calc(16cm - ' + thickness + ' + ' + pixel + ')', 'calc(12cm + ' + pixel + ')', thickness, length); // vertical (finally visible)

  // --- MIDPOINTS ---

  // Top-center
  drawLine('calc(8.5cm - 0.1mm)', 'calc(1cm - ' + length + ' - ' + pixel + ')', thickness, length);

  // Bottom-center
  drawLine('calc(8.5cm - 0.1mm)', 'calc(12cm + ' + pixel + ')', thickness, length);

  // Left-center
  drawLine('calc(1cm - ' + length + ' - ' + pixel + ')', 'calc(6.5cm - 0.1mm)', length, thickness);

  // Right-center
  drawLine('calc(16cm + ' + pixel + ')', 'calc(6.5cm - 0.1mm)', length, thickness);
};
