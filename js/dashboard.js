var ctx = document.getElementById('diagnosis_chart').getContext('2d');
var myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"],
        datasets: [{
            label: 'Diagnosis',
            backgroundColor: '#bbdefb',
            borderColor: window.theme.primary,
            hoverBackgroundColor: window.theme.primary,
            hoverBorderColor: window.theme.primary,
            data: [13, 2, 5, 7, 11, 1, 21, 4, 3, 4, 4, 12],
            barPercentage: .75,
            categoryPercentage: .5
        }]
    },
    options: {
        scales: {
            yAxes: [{
                ticks: {
                    beginAtZero: true
                }
            }]
        }
    }
});




var ctx = document.getElementById('procedure_chart').getContext('2d');
var myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: ["Appen", "Cat", "CS", "CT", "CHD", "CABG", "MRI", "X-RAY", "US", "EKG", "EMG", "EEG"],
        datasets: [{
            label: 'Procedure',
            backgroundColor: '#9fa8da',
            borderColor: window.theme.primary,
            hoverBackgroundColor: window.theme.primary,
            hoverBorderColor: window.theme.primary,
            data: [4, 6, 1, 1, 11, 6, 21, 7, 10, 13, 1, 3],
            barPercentage: .75,
            categoryPercentage: .5
            }]
    },
    options: {
        scales: {
            yAxes: [{
                ticks: {
                    beginAtZero: true
                }
            }]
        }
    }
});






